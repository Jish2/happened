import { getGreetingByName } from "@/gen/openapi";
import { AxiosError } from "axios";
import { Link } from "expo-router";
import { useEffect, useState } from "react";
import { Text, TouchableOpacity, View } from "react-native";

export default function Page() {
  const { user } = useUser();

  const [name, setName] = useState<string>("");

  useEffect(() => {
    const loadName = async () => {
      try {
        const res = await getGreetingByName("name");
        setName(res.data.message);
      } catch (e) {
        if (e instanceof AxiosError) console.log(e.message);
      }
    };

    loadName();
  }, []);

  return (
    <View className="justify-center items-center flex-1">
      <Text>Name: {name || "no name"}</Text>
      <SignedIn>
        <Text>Hello {user?.emailAddresses[0].emailAddress}</Text>
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
