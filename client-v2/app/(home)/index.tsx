import { getGreeting } from "@/gen/openapi";
import { SignedIn, SignedOut, useClerk, useUser } from "@clerk/clerk-expo";
import axios, { AxiosError } from "axios";
import { Link } from "expo-router";
import { useEffect, useState } from "react";
import { Text, TouchableOpacity, View } from "react-native";

axios.defaults.baseURL =
  process.env.EXPO_PUBLIC_API_URL || "http://localhost:8080";

export default function Page() {
  const { user } = useUser();
  const { signOut } = useClerk();

  const [name, setName] = useState<string>("");

  useEffect(() => {
    const loadName = async () => {
      try {
        const res = await getGreeting("name");
        setName(res.data.message);
      } catch (e) {
        if (e instanceof AxiosError) console.error(e.message);
      }
    };

    loadName();
  }, []);

  return (
    <View className="justify-center items-center flex-1">
      <Text>Name: {name || "no name"}</Text>
      <SignedIn>
        <Text>Hello {user?.phoneNumbers[0]?.phoneNumber}</Text>
        <Text>{JSON.stringify(user, null, 2)}</Text>
        <TouchableOpacity
          onPress={() => {
            signOut();
          }}
        >
          <Text>Sign Out</Text>
        </TouchableOpacity>
      </SignedIn>
      <SignedOut>
        <Link href="/(auth)/sign-in">
          <Text>Sign In</Text>
        </Link>
        <Link href="/(auth)/sign-up">
          <Text>Sign Up</Text>
        </Link>
      </SignedOut>
    </View>
  );
}
