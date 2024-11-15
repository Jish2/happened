import {
  Image,
  StyleSheet,
  Platform,
  Text,
  TouchableOpacity,
} from "react-native";
import { HelloWave } from "@/components/HelloWave";
import ParallaxScrollView from "@/components/ParallaxScrollView";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { SignedIn, SignedOut, useClerk, useUser } from "@clerk/clerk-expo";
import { Link } from "expo-router";
import { useEffect, useState } from "react";
import {
  GreetRequestSchema,
  HappenedService,
} from "@/gen/protos/v1/happened_service_pb";
import { ConnectError, createClient } from "@connectrpc/connect";
import { createXHRGrpcWebTransport } from "@/app/custom-transport";
// Needed to polyfill TextEncoder/ TextDecoder
import "fast-text-encoding";
import { create } from "@bufbuild/protobuf";
import { polyfills } from "@/app/polyfill.native";

polyfills();

export default function HomeScreen() {
  const [greeting, setGreeting] = useState("none");

  const client = createClient(
    HappenedService,
    createXHRGrpcWebTransport({
      baseUrl: "http://localhost:8080",
    }),
  );
  useEffect(() => {
    const getData = async () => {
      try {
        const request = create(GreetRequestSchema, {
          name: "",
        });

        const response = await client.greet(request);
        console.log("response", response);
        setGreeting(response.greeting);
      } catch (e) {
        if (e instanceof ConnectError) {
          console.error("error calling greet", String(e));
        }
      }
    };
    getData();
  }, [client]);

  const { user } = useUser();
  const { signOut } = useClerk();

  return (
    <ParallaxScrollView
      headerBackgroundColor={{ light: "#A1CEDC", dark: "#1D3D47" }}
      headerImage={
        <Image
          source={require("@/assets/images/partial-react-logo.png")}
          style={styles.reactLogo}
        />
      }
    >
      <ThemedView style={styles.titleContainer}>
        <ThemedText type="title">Welcome {greeting}!</ThemedText>
        <HelloWave />
        <Text className="bg-blue-500 text-white">Baz</Text>
      </ThemedView>
      <ThemedView style={styles.titleContainer}>
        <SignedIn>
          <Text>Hello {JSON.stringify(user?.phoneNumbers[0].phoneNumber)}</Text>
          <TouchableOpacity onPress={() => signOut()}>
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
      </ThemedView>
      <ThemedView style={styles.stepContainer}>
        <ThemedText type="subtitle">Step 1: Try it</ThemedText>
        <ThemedText>
          Edit{" "}
          <ThemedText type="defaultSemiBold">app/(tabs)/index.tsx</ThemedText>{" "}
          to see changes. Press{" "}
          <ThemedText type="defaultSemiBold">
            {Platform.select({ ios: "cmd + d", android: "cmd + m" })}
          </ThemedText>{" "}
          to open developer tools.
        </ThemedText>
      </ThemedView>
      <ThemedView style={styles.stepContainer}>
        <ThemedText type="subtitle">Step 2: Explore</ThemedText>
        <ThemedText>
          Tap the Explore tab to learn more about what's included in this
          starter app.
        </ThemedText>
      </ThemedView>
      <ThemedView style={styles.stepContainer}>
        <ThemedText type="subtitle">Step 3: Get a fresh start</ThemedText>
        <ThemedText>
          When you're ready, run{" "}
          <ThemedText type="defaultSemiBold">npm run reset-project</ThemedText>{" "}
          to get a fresh <ThemedText type="defaultSemiBold">app</ThemedText>{" "}
          directory. This will move the current{" "}
          <ThemedText type="defaultSemiBold">app</ThemedText> to{" "}
          <ThemedText type="defaultSemiBold">app-example</ThemedText>.
        </ThemedText>
      </ThemedView>
    </ParallaxScrollView>
  );
}

const styles = StyleSheet.create({
  titleContainer: {
    flexDirection: "row",
    alignItems: "center",
    gap: 8,
  },
  stepContainer: {
    gap: 8,
    marginBottom: 8,
  },
  reactLogo: {
    height: 178,
    width: 290,
    bottom: 0,
    left: 0,
    position: "absolute",
  },
});
