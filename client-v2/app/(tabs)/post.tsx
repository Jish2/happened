import { SafeAreaView } from "react-native-safe-area-context";
import { Text, TouchableOpacity, Image } from "react-native";
import { useState } from "react";
import * as ImagePicker from "expo-image-picker";
import { ImagePickerAsset } from "expo-image-picker";
import * as Crypto from "expo-crypto";

export default function PostTab() {
  const [images, setImages] = useState<ImagePickerAsset[]>([]);

  const pickImage = async () => {
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ["images"],
      allowsEditing: false,
      aspect: [4, 3],
      quality: 1,
      allowsMultipleSelection: true,
      selectionLimit: 10,
      exif: true,
    });
    console.log(result);
    if (!result.canceled) {
      setImages(result.assets);
    }
  };

  return (
    <SafeAreaView className="flex items-center justify-center bg-white h-full w-full">
      <Text>Hello Post Tab</Text>

      <TouchableOpacity onPress={pickImage}>
        <Text>Select Image for Upload</Text>
      </TouchableOpacity>
      {images && (
        <Image
          source={{ uri: images[0]?.uri }}
          className="aspect-square w-1/2 h-1/2"
        />
      )}

      <TouchableOpacity
        onPress={async () => {
          // try {
          //   const imageKey = Crypto.randomUUID();
          //   console.log("imageKey", imageKey);
          //   // const { uploadUrl, headers, method } =
          //     // await client.getUploadImageURL({
          //     //   imageKey,
          //     // });
          //   // console.log("messages", images[0].base64);
          //   // const res = await fetch(uploadUrl, {
          //   //   headers: headers,
          //   //   method
          //   // })
          //   // setUrl(uploadUrl);
          //   console.log("uploadUrl", uploadUrl);
          // } catch (e) {
          //   if (e instanceof ConnectError) {
          //     console.error(e.cause, e.details, e.code);
          //   }
          // }
        }}
      >
        <Text>Upload Image</Text>
      </TouchableOpacity>
    </SafeAreaView>
  );
}
