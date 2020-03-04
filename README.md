Start Here
==================
[Start here](https://github.com/andrewarrow/feedbacks/blob/master/README.md)

Read [the feedbacks readme](https://github.com/andrewarrow/feedbacks/blob/master/README.md) first. 


ANDROID_NDK_HOME=~/android-ndk-r21/ gomobile build -target=android github.com/andrewarrow/feedback/mobile
ANDROID_NDK_HOME=~/android-ndk-r21/ gomobile install github.com/andrewarrow/feedback/mobile

cd Android/Sdk/emulator
./emulator -avd Galaxy_Nexus_API_23

cd Android/Sdk/platform-tools/
./adb install ~/src/feedback/mobile.apk
