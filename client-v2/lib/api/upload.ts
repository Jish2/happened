import axios from "axios";
import * as FileSystem from "expo-file-system";
import { Buffer } from "buffer";

export async function uploadImage(
  upload_url: string,
  signed_headers: Record<string, string>,
  imageUri: string,
) {
  const base64 = await FileSystem.readAsStringAsync(imageUri, {
    encoding: FileSystem.EncodingType.Base64,
  });
  console.log("base64", base64.length);

  try {
    const bytes = new Uint8Array(Buffer.from(base64, "base64"));
    console.log("uploading image");
    console.log("uploading image");
    const resp = await axios.put(upload_url, bytes, {
      headers: signed_headers,
    });

    console.log("done uploading image", resp);
  } catch (e) {
    console.error(e);
  }
}
