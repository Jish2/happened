import { SafeAreaView } from "react-native-safe-area-context";
import { Text, TouchableOpacity, Image } from "react-native";
import { useState } from "react";
import * as ImagePicker from "expo-image-picker";
import { ImagePickerAsset } from "expo-image-picker";
import * as Crypto from "expo-crypto";
import { postCreateUploadUrl } from "@/lib/api/happened";
import { uploadImage } from "@/lib/api/upload";

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
    // console.log(result);
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
          const imageKey = Crypto.randomUUID();
          const res = await postCreateUploadUrl({ image_key: imageKey }).catch(
            console.error,
          );

          if (res) {
            console.log("got response");

            const { method, signed_headers, upload_url } = res.data;

            console.log("data", res.data);

            await uploadImage(upload_url, signed_headers, images[0].uri);
          }
        }}
      >
        <Text>Upload Image</Text>
      </TouchableOpacity>
    </SafeAreaView>
  );
}
