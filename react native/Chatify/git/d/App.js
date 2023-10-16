import { useCallback, useLayoutEffect, useState } from "react";
import * as SplashScreen from "expo-splash-screen";
import * as Font from "expo-font";
import { LogBox, SafeAreaView } from "react-native";
import { Provider } from "react-redux";

import { store } from "./store/store";
import Routes from "./navigation/Routes";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { SafeAreaProvider } from "react-native-safe-area-context";

LogBox.ignoreLogs(["AsyncStorage has been extracted"]);

//AsyncStorage.clear(); // force logging out

SplashScreen.preventAutoHideAsync();

const App = () => {
  const [isAppLoaded, setIsAppLoaded] = useState(false);

  useState(() => {
    // loading app setups
    const prepare = async () => {
      try {
        await Font.loadAsync({
          "Raleway-Regular": require("./assets/fonts/Raleway-Regular.ttf"),
          "Raleway-ExtraBold": require("./assets/fonts/Raleway-ExtraBold.ttf"),
        });
      } catch (error) {
        console.log.error();
      } finally {
        setIsAppLoaded(true);
      }
    };
    prepare();
  }, []);

  const onLayout = useCallback(async () => {
    if (isAppLoaded) {
      await SplashScreen.hideAsync();
    }
  });

  if (!isAppLoaded) {
    return null;
  }

  return (
    <Provider store={store}>
      <SafeAreaProvider onLayout={onLayout}>
        <Routes />
      </SafeAreaProvider>
    </Provider>
  );
};

export default App;
