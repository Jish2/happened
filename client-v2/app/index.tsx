import { Text, View } from "react-native";
import axios, { AxiosError } from "axios";
import { getGreetingByName } from "@/gen/openapi";
import { useEffect, useState } from "react";

import "@/global.css";

axios.defaults.baseURL = "http://localhost:8080";

export default function Index() {
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
    <View
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Text className={"bg-red-500"}>Name: {name || "no name"}</Text>
      <Text>Edit app/index.tsx to edit this screen.</Text>
    </View>
  );
}
