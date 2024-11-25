import { View } from "react-native";
import { GooglePlacesAutocomplete } from "react-native-google-places-autocomplete";

export const SearchAutocomplete = () => {
  return (
    <>
      <View className="p-4 w-screen h-full absolute bg-white">
        <GooglePlacesAutocomplete
          placeholder="Search"
          onPress={(data, details = null) => {
            // 'details' is provided when fetchDetails = true
            console.log(data, details);
          }}
          query={{
            key: process.env.EXPO_PUBLIC_GOOGLE_PLACES_API_KEY,
            language: "en",
          }}
          enablePoweredByContainer={false}
          // GooglePlacesDetailsQuery={{}}
        />
      </View>
    </>
  );
};
