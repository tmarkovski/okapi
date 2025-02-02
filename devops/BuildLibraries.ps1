param
(
    [ValidateSet('Windows', 'Windows-GNU', 'MacOS', 'Linux', 'Linux-ARM', 'iOS', 'Android-x86', 'Android-ARM')]
    $Platform,
    $OutLocation,
    $AndroidNdkHome
)

$ErrorActionPreference = "Stop"

# Ensure target directory exists and gather invocation info
$InvocationPath = (Get-Item .).FullName
$TargetOutput = $ExecutionContext.SessionState.Path.GetUnresolvedProviderPathFromPSPath([IO.Path]::Combine($InvocationPath, $OutLocation))
mkdir -p $TargetOutput

if ($null -eq $Platform) { throw "Parameter -Platform must be specified." }
if ($null -eq $OutLocation) { throw "Parameter -OutLocation must be specified." }

# Allows to execute this script from any context
Set-Location $PSScriptRoot/../native

try {
    switch ($Platform) {
        Windows {
            rustup target add x86_64-pc-windows-msvc
            cargo build --release --target x86_64-pc-windows-msvc

            Copy-Item -Path .\target\x86_64-pc-windows-msvc\release\okapi.dll -Destination $TargetOutput
            Copy-Item -Path .\target\x86_64-pc-windows-msvc\release\okapi.lib -Destination $TargetOutput
            break
        }
        Windows-GNU {
            rustup target add x86_64-pc-windows-gnu
            cargo build --release --target x86_64-pc-windows-gnu

            Copy-Item -Path .\target\x86_64-pc-windows-gnu\release\okapi.dll -Destination $TargetOutput
            Copy-Item -Path .\target\x86_64-pc-windows-gnu\release\libokapi.dll.a -Destination $TargetOutput
            break
        }
        Linux {
            rustup target add x86_64-unknown-linux-gnu
            cargo build --release --target x86_64-unknown-linux-gnu

            Copy-Item -Path .\target\x86_64-unknown-linux-gnu\release\libokapi.so -Destination $TargetOutput
            Copy-Item -Path .\target\x86_64-unknown-linux-gnu\release\libokapi.a -Destination $TargetOutput
            break
        }
        Linux-ARM {
            sudo apt-get install gcc-arm-linux-gnueabihf gcc-aarch64-linux-gnu
            rustup target add armv7-unknown-linux-gnueabihf aarch64-unknown-linux-gnu
            cargo build --release --target armv7-unknown-linux-gnueabihf
            cargo build --release --target aarch64-unknown-linux-gnu

            mkdir -p $TargetOutput/linux-armv7/
            mkdir -p $TargetOutput/linux-aarch64/
            Copy-Item -Path .\target\armv7-unknown-linux-gnueabihf\release\libokapi.so -Destination "$TargetOutput/linux-armv7"
            Copy-Item -Path .\target\armv7-unknown-linux-gnueabihf\release\libokapi.a -Destination "$TargetOutput/linux-armv7"
            Copy-Item -Path .\target\aarch64-unknown-linux-gnu\release\libokapi.so -Destination "$TargetOutput/linux-aarch64"
            Copy-Item -Path .\target\aarch64-unknown-linux-gnu\release\libokapi.a -Destination "$TargetOutput/linux-aarch64"
            break
        }
        MacOS {
            # Need OSX11 for the ARM64 libraries
            $env:SDKROOT=$(xcrun -sdk macosx11.1 --show-sdk-path)
            $env:MACOSX_DEPLOYMENT_TARGET=$(xcrun -sdk macosx11.1 --show-sdk-platform-version)

            rustup target add x86_64-apple-darwin aarch64-apple-darwin
            cargo build --release --target x86_64-apple-darwin
            cargo build --release --target aarch64-apple-darwin
            # Create the fat binaries.
            lipo -create "./target/x86_64-apple-darwin/release/libokapi.a" "./target/aarch64-apple-darwin/release/libokapi.a" -output "$TargetOutput/libokapi.a"
            lipo -create "./target/x86_64-apple-darwin/release/libokapi.dylib" "./target/aarch64-apple-darwin/release/libokapi.dylib" -output "$TargetOutput/libokapi.dylib"
            break
        }
        iOS {
            # Need OSX11 for the ARM64 libraries
            $env:SDKROOT=$(xcrun -sdk macosx11.1 --show-sdk-path)
            $env:MACOSX_DEPLOYMENT_TARGET=$(xcrun -sdk macosx11.1 --show-sdk-platform-version)

            # Install nightly to allow for building ios sim.
            rustup toolchain install nightly --allow-downgrade
            rustup target add x86_64-apple-ios aarch64-apple-ios
            rustup target add aarch64-apple-ios-sim --toolchain nightly
            
            cargo +nightly build --release --target aarch64-apple-ios-sim
            cargo build --release --target x86_64-apple-ios
            cargo build --release --target aarch64-apple-ios

            # Create the fat binaries, cargo-lipo doesn't support ios sim aarch64
            lipo -create "./target/x86_64-apple-ios/release/libokapi.a" "./target/aarch64-apple-ios-sim/release/libokapi.a" -output "$TargetOutput/libokapi_simulator.a"
            Copy-Item -Path "./target/aarch64-apple-ios/release/libokapi.a" -Destination "$TargetOutput/libokapi.a"
            break
        }
        Android-ARM {
            cargo install cargo-ndk
            rustup target add aarch64-linux-android armv7-linux-androideabi
            cargo ndk --target armv7-linux-androideabi --target aarch64-linux-android -o "$TargetOutput" build --release
            break
        }
        Android-x86 {
            cargo install cargo-ndk
            rustup target add i686-linux-android x86_64-linux-android
            cargo ndk --target x86_64-linux-android --target i686-linux-android -o "$TargetOutput" build --release
            break
        }
    }
}
finally {
    # Return to invocation context
    Set-Location $InvocationPath
}