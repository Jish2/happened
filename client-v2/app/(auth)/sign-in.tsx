import { useSignIn } from "@clerk/clerk-expo";
import { Link, useRouter } from "expo-router";
import {
  Text,
  TextInput,
  View,
  TouchableOpacity,
  TouchableHighlight,
} from "react-native";
import { useCallback, useState } from "react";
import { Ionicons } from "@expo/vector-icons";
import { SignInWithOAuth } from "@/components/sign-in-with-o-auth";

export default function Page() {
  const { signIn, setActive, isLoaded } = useSignIn();
  const router = useRouter();

  const [pendingVerification, setPendingVerification] = useState(false);
  const [phoneNumber, setPhoneNumber] = useState<string>("");
  const [code, setCode] = useState("");

  const onSignInPress = useCallback(async () => {
    if (!isLoaded) {
      return;
    }

    try {
      const { supportedFirstFactors } = await signIn.create({
        identifier: phoneNumber,
      });

      const phoneNumberId = supportedFirstFactors?.find(
        (factor) => factor.strategy === "phone_code",
      )?.phoneNumberId;

      if (!phoneNumberId) throw new Error("No phoneNumberId");

      await signIn.prepareFirstFactor({
        strategy: "phone_code",
        phoneNumberId,
      });

      setPendingVerification(true);
    } catch (err: unknown) {
      // See https://clerk.com/docs/custom-flows/error-handling
      // for more info on error handling
      console.error(JSON.stringify(err, null, 2));
    }

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isLoaded, phoneNumber]);

  const onPressVerify = useCallback(async () => {
    if (!isLoaded) {
      return;
    }

    try {
      const completeSignIn = await signIn.attemptFirstFactor({
        strategy: "phone_code",
        code,
      });

      if (completeSignIn.status === "complete") {
        await setActive({ session: completeSignIn.createdSessionId });
        router.replace("/(home)");
      } else {
        // See https://clerk.com/docs/custom-flows/error-handling
        // for more info on error handling
        console.error(JSON.stringify(completeSignIn, null, 2));
      }
    } catch (err: unknown) {
      console.error(JSON.stringify(err, null, 2));
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [code, isLoaded]);

  return (
    <View className="h-full justify-center pb-16">
      <View className="absolute left-4 top-8">
        <TouchableHighlight>
          <Link href="/(home)">
            <Ionicons name="chevron-back" size={32} color="black" />
          </Link>
        </TouchableHighlight>
      </View>
      <View className="p-6">
        <Text className="text-7xl">login to your account.</Text>
      </View>

      <>
        {pendingVerification ? (
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
              className="px-3 h-12 flex-row gap-2 items-center justify-center rounded-lg border-2 bg-black"
            >
              <View className="flex flex-row items-center gap-2">
                <Text className="text-white">Verify</Text>
              </View>
            </TouchableOpacity>
          </View>
        ) : (
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
        )}
      </>
    </View>
  );
}
