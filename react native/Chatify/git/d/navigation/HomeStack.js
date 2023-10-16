import { View, Text } from "react-native";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import MaterialCommunityIcons from "@expo/vector-icons/MaterialCommunityIcons";
import { Feather } from "@expo/vector-icons";

import ChatListScreen from "../Screens/ChatListScreen";
import SettingsScreen from "../Screens/SettingsScreen";

import { COLORS } from "../constants";

const Tab = createBottomTabNavigator();

const HomeStack = () => {
  return (
    <Tab.Navigator>
      <Tab.Screen
        name="ChatList"
        component={ChatListScreen}
        options={({ route }) => ({
          headerShown: false,
          tabBarIcon: (props) => {
            return (
              <MaterialCommunityIcons
                name="chat-outline"
                {...props}
                style={{
                  color: props.focused ? COLORS.mainOrange : COLORS.notFocused,
                }}
              />
            );
          },
          tabBarLabel: ({ focused }) => {
            return (
              <View>
                <Text
                  style={{
                    color: focused ? COLORS.mainOrange : COLORS.notFocused,
                  }}
                >
                  Chats
                </Text>
              </View>
            );
          },
        })}
      />

      <Tab.Screen
        name="Settings"
        component={SettingsScreen}
        options={{
          headerShown: false,
          tabBarIcon: (props) => {
            return (
              <Feather
                name="settings"
                {...props}
                style={{
                  color: props.focused ? COLORS.mainOrange : COLORS.notFocused,
                }}
              />
            );
          },
          tabBarLabel: ({ focused }) => {
            return (
              <View>
                <Text
                  style={{
                    color: focused ? COLORS.mainOrange : COLORS.notFocused,
                  }}
                >
                  Settings
                </Text>
              </View>
            );
          },
        }}
      />
    </Tab.Navigator>
  );
};

export default HomeStack;
