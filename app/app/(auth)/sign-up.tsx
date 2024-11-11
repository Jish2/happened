import * as React from 'react';
import { TextInput, View, Pressable, Text } from 'react-native';
import { useSignUp } from '@clerk/clerk-expo';
import { Link, useRouter } from 'expo-router';
import SignInWithOAuth from '@/components/SignInWithOAuth';
import { Ionicons } from '@expo/vector-icons';
import { TouchableOpacity } from 'react-native';

export default function SignUpScreen() {
  const { isLoaded, signUp, setActive } = useSignUp();
  const router = useRouter();

  const [phoneNumber, setPhoneNumber] = React.useState('');
  const [pendingVerification, setPendingVerification] = React.useState(false);
  const [code, setCode] = React.useState('');

  const onSignUpPress = async () => {
    if (!isLoaded) {
      return;
    }

    try {
      await signUp.create({
        phoneNumber,
      });

      await signUp.preparePhoneNumberVerification({ strategy: 'phone_code' });

      // await signUp.prepareEmailAddressVerification({ strategy: "email_code" });

      setPendingVerification(true);
    } catch (err: any) {
      // See https://clerk.com/docs/custom-flows/error-handling
      // for more info on error handling
      console.error(JSON.stringify(err, null, 2));
    }
  };

  const onPressVerify = async () => {
    if (!isLoaded) {
      return;
    }

    try {
      const completeSignUp = await signUp.attemptPhoneNumberVerification({
        code,
      });

      if (completeSignUp.status === 'complete') {
        await setActive({ session: completeSignUp.createdSessionId });
        router.replace('/');
      } else {
        console.error(JSON.stringify(completeSignUp, null, 2));
      }
    } catch (err: any) {
      // See https://clerk.com/docs/custom-flows/error-handling
      // for more info on error handling
      console.error(JSON.stringify(err, null, 2));
    }
  };

  return (
    <View className="h-full justify-center pb-16">
      <View className="absolute left-4 top-32">
        <Pressable>
          <Link href="/">
            <Ionicons name="chevron-back" size={32} color="black" />
          </Link>
        </Pressable>
      </View>
      <View className="p-6">
        <Text className="text-7xl">create an account.</Text>
      </View>
      {!pendingVerification && (
        <>
          <View className="p-6 pb-0 gap-2">
            <TextInput
              autoCapitalize="none"
              value={phoneNumber}
              keyboardType="phone-pad"
              placeholder="Phone Number"
              onChangeText={(text) => setPhoneNumber(text)}
              className="bg-neutral-100 dark:bg-neutral-900 border-2 border-black/20 dark:border-white/20 rounded-lg h-12 px-4 text-neutral-950 dark:text-neutral-50"
            />
            <TouchableOpacity
              onPress={onSignUpPress}
              className="px-3 h-12 flex-row gap-2 items-center justify-center rounded-lg border-2 bg-black">
              <View className="flex flex-row items-center gap-2">
                <Text className="text-white">Continue</Text>
              </View>
            </TouchableOpacity>
          </View>
          <SignInWithOAuth />
        </>
      )}
      {pendingVerification && (
        <View className="p-6 pb-0 gap-2">
          <TextInput
            autoCapitalize="none"
            value={code}
            keyboardType="phone-pad"
            placeholder="verification code"
            onChangeText={(text) => setCode(text)}
            className="bg-neutral-100 dark:bg-neutral-900 border-2 border-black/20 dark:border-white/20 rounded-lg h-12 px-4 text-neutral-950 dark:text-neutral-50"
          />

          <TouchableOpacity
            onPress={onPressVerify}
            className="px-3 h-12 flex-row gap-2 items-center justify-center rounded-lg border-2 bg-black">
            <View className="flex flex-row items-center gap-2">
              <Text className="text-white">Verify</Text>
            </View>
          </TouchableOpacity>
        </View>
      )}
    </View>
  );
}
