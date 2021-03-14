Start Here
==================
[Start here](https://github.com/andrewarrow/feedbacks/blob/master/README.md)

Read [the feedbacks readme](https://github.com/andrewarrow/feedbacks/blob/master/README.md) first. 


  brew install mariadb
  brew services start mariadb









ANDROID_NDK_HOME=~/android-ndk-r21/ gomobile build -target=android github.com/andrewarrow/feedback/mobile
ANDROID_NDK_HOME=~/android-ndk-r21/ gomobile install github.com/andrewarrow/feedback/mobile

// detectTeamID in 
gp/pkg/mod/golang.org/x/mobile@v0.0.0-20200222142934-3c8601c510d0/cmd/gomobile/build_iosapp.go 
gomobile build -ldflags=-extldflags=-Wl,-z,relro -target=ios -bundleid=x github.com/andrewarrow/feedback/mobile

cd Android/Sdk/emulator
./emulator -avd Galaxy_Nexus_API_23

cd Android/Sdk/platform-tools/
./adb install ~/src/feedback/mobile.apk

./sdkmanager  --sdk_root=.. --licenses
./sdkmanager  --sdk_root=.. --install 'emulator'
./sdkmanager  --sdk_root=.. --install 'system-images;android-23;google_apis;x86_64'

https://github.com/benvium/libimobiledevice-macosx

./avdmanager create avd --force --name testAVD --abi google_apis/x86_64 --package 'system-images;android-23;google_apis;x86_64'

 vi ~/.android/avd/generic_10.avd/config.ini
