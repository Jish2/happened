import { SafeAreaView } from "react-native-safe-area-context";
import { Text, TouchableOpacity, Image } from "react-native";
import { useState } from "react";
import * as ImagePicker from "expo-image-picker";
import { ImagePickerAsset } from "expo-image-picker";

import { uploadImages } from "@/lib/api/upload";
import { useMutation } from "@tanstack/react-query";

export default function PostTab() {
  const [images, setImages] = useState<ImagePickerAsset[]>([]);
  // Access the client
  // const queryClient = useQueryClient();

  // Queries
  const { isPending: isUploadPending, mutateAsync: uploadImagesMutation } =
    useMutation({
      mutationFn: uploadImages,
      onSuccess: () => console.log("successfully uploaded images"),
      onError: (e) => console.error(e),
    });
  // uploadImagesMutation.mutate

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
          uploadImagesMutation(images);
        }}
      >
        <Text>Upload Image</Text>
      </TouchableOpacity>
      {isUploadPending && <Text>Uploading {images.length} images...</Text>}
    </SafeAreaView>
  );
}
