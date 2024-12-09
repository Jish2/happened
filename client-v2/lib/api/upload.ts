import axios from "axios";
import * as FileSystem from "expo-file-system";
import { Buffer } from "buffer";
import * as Crypto from "expo-crypto";
import { postCreateUploadUrl } from "@/lib/api/happened";
import { ImagePickerAsset } from "expo-image-picker";
import { ImagePropsAndroid } from "react-native";

export async function uploadImage(image: ImagePickerAsset) {
  const imageKey = Crypto.randomUUID();
  const res = await postCreateUploadUrl({ image_key: imageKey }).catch(
    console.error,
  );
  if (!res) {
    console.error("no response from post create upload url");
    return;
  }

  console.log("got response");
  const { method, signed_headers, upload_url } = res.data;

  const base64 = await FileSystem.readAsStringAsync(image.uri, {
    encoding: FileSystem.EncodingType.Base64,
  });
  console.log("base64", base64.length);

  try {
    const bytes = new Uint8Array(Buffer.from(base64, "base64"));
    console.log("uploading image");
    const resp = await axios.put(upload_url, bytes, {
      headers: signed_headers,
    });

    console.log("done uploading image", resp);
  } catch (e) {
    console.error(e);
  }
}

export async function uploadImages(images: ImagePickerAsset[]) {
  const promises = [];
  for (const image of images) {
    const uploadPromise = uploadImage(image);
    promises.push(uploadPromise);
  }

  await Promise.all(promises);
  console.log("finished uploading all images");
}
