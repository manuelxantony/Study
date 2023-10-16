import { createStackNavigator } from "@react-navigation/stack";
import { Button } from "react-native";

import HomeStack from "./HomeStack";
import ChatScreen from "../Screens/ChatScreen";
import NewChatScreen from "../Screens/NewChatScreen";

const Stack = createStackNavigator();

const MainStack = () => {
  return (
    <Stack.Navigator>
      <Stack.Group>
        <Stack.Screen
          name="Home"
          component={HomeStack}
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="ChatScreen"
          component={ChatScreen}
          options={{
            headerShown: false,
          }}
        />
      </Stack.Group>
      <Stack.Group
        screenOptions={({ navigation }) => ({
          presentation: "modal",
          headerLeft: () => (
            <Button title="Cancel" onPress={navigation.goBack} />
          ),
        })}
      >
        <Stack.Screen
          name="NewChat"
          component={NewChatScreen}
          options={{
            headerShown: false,
          }}
        />
      </Stack.Group>
    </Stack.Navigator>
  );
};

export default MainStack;
