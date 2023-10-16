import { StatusBar } from "expo-status-bar";
import { useCallback, useEffect, useState } from "react";
import { SafeAreaView, StyleSheet, Text, View, Button } from "react-native";

const Total = ({ getTo }) => {
  const [total, setTotal] = useState(0);

  useEffect(() => {
    console.log("calculating");
    setTotal(getTo());
  }, [getTo]);

  return (
    <View>
      <Text>Total: {total}</Text>
    </View>
  );
};

export default Total;

const styles = StyleSheet.create({});
