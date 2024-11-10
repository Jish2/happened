import { useSignIn } from "@clerk/clerk-expo";
import { Link, useRouter } from "expo-router";
import {
  Text,
  TextInput,
  View,
  TouchableOpacity,
  TouchableHighlight,
} from "react-native";
import React from "react";
import { Ionicons } from "@expo/vector-icons";
import SignInWithOAuth from "@/components/SignInWithOAuth";

export default function Page() {
  const { signIn, setActive, isLoaded } = useSignIn();
  const router = useRouter();

  const [phoneNumber, setPhoneNumber] = React.useState("");

  const onSignInPress = React.useCallback(async () => {
    if (!isLoaded) {
      return;
    }

    try {
      const signInAttempt = await signIn.create({
        identifier: phoneNumber,
      });

      if (signInAttempt.status === "complete") {
        await setActive({ session: signInAttempt.createdSessionId });
        router.replace("/");
      } else {
        // See https://clerk.com/docs/custom-flows/error-handling
        // for more info on error handling
        console.error(JSON.stringify(signInAttempt, null, 2));
      }
    } catch (err: any) {
      console.error(JSON.stringify(err, null, 2));
    }
  }, [isLoaded, phoneNumber, router, setActive, signIn]);

  return (
    <View className="h-full justify-center pb-16">
      <View className="absolute left-4 top-32">
        <TouchableHighlight>
          <Link href="/">
            <Ionicons name="chevron-back" size={32} color="black" />
          </Link>
        </TouchableHighlight>
      </View>
      <View className="p-6">
        <Text className="text-7xl">login to your account.</Text>
      </View>

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
            onPress={onSignInPress}
            className="px-3 h-12 flex-row gap-2 items-center justify-center rounded-lg border-2 bg-black"
          >
            <View className="flex flex-row items-center gap-2">
              <Text className="text-white">Continue</Text>
            </View>
          </TouchableOpacity>
        </View>
        <SignInWithOAuth />
      </>
    </View>
  );
}
