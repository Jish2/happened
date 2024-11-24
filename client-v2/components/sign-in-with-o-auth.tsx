import { useCallback, useEffect } from "react";
import * as WebBrowser from "expo-web-browser";
import { Text, View, TouchableOpacity } from "react-native";
import { useRouter } from "expo-router";
import { useOAuth } from "@clerk/clerk-expo";
import * as Linking from "expo-linking";
import { AntDesign } from "@expo/vector-icons";

export const useWarmUpBrowser = () => {
  useEffect(() => {
    // Warm up the android browser to improve UX
    // https://docs.expo.dev/guides/authentication/#improving-user-experience
    void WebBrowser.warmUpAsync();
    return () => {
      void WebBrowser.coolDownAsync();
    };
  }, []);
};

WebBrowser.maybeCompleteAuthSession();

export const SignInWithOAuth = () => {
  useWarmUpBrowser();

  const router = useRouter();

  const { startOAuthFlow } = useOAuth({ strategy: "oauth_google" });

  const onPress = useCallback(async () => {
    try {
      const { createdSessionId, setActive } = await startOAuthFlow({
        redirectUrl: Linking.createURL("/(auth)", { scheme: "happened" }),
      });

      if (createdSessionId) {
        setActive!({ session: createdSessionId });
      } else {
        // Use signIn or signUp for next steps such as MFA
        console.log("session created but no further action taken");
      }
      router.navigate("/(home)");
    } catch (err) {
      console.error("OAuth error", err);
    }
  }, [router, startOAuthFlow]);

  return (
    <View>
      <View className="flex-row justify-center gap-2 p-6">
        <TouchableOpacity
          onPress={onPress}
          className="h-10 px-4 flex-row gap-2 items-center justify-center rounded-lg border-2 w-1/2"
        >
          <View className="flex flex-row items-center gap-2">
            <AntDesign name="google" size={16} color="black" />
            <Text>Sign In with Google</Text>
          </View>
        </TouchableOpacity>
        <TouchableOpacity
          onPress={() => console.log("Pressed")}
          className="h-10 px-4 flex-row gap-2 items-center justify-center rounded-lg border-2 flex-1 opacity-50"
          disabled
        >
          <View className="flex flex-row items-center gap-2">
            <AntDesign name="apple1" size={16} color="black" />
            <Text>Sign In with Apple</Text>
          </View>
        </TouchableOpacity>
      </View>
    </View>
  );
};
