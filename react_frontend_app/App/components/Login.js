import React from "react";
import {
  StyleSheet,
  Text,
  View,
  TextInput,
  TouchableOpacity,
} from "react-native";

export default function Login({ navigation }) {
  const registerButtonHandler = () => {
    navigation.navigate("Registration");
  };

  return (
    <View style={styles.logIn}>
      <Text style={styles.title}>Hello.</Text>
      <Text style={styles.title}>Welcome Back</Text>

      <View style={styles.mainFormat}>
        <TextInput
          style={styles.textInput}
          placeholder="User-Name"
          underlineColorAndroid={"transparent"}
        />
        <TextInput
          style={styles.textInput}
          placeholder="Enter Password"
          secureTextEntry={true}
          underlineColorAndroid={"transparent"}
        />
      </View>
      <TouchableOpacity style={styles.button}>
        <Text style={styles.buttonText}>Log In</Text>
      </TouchableOpacity>

      <TouchableOpacity onPress={registerButtonHandler} style={styles.button}>
        <Text style={styles.buttonText}>Register</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  logIn: {
    alignSelf: "stretch",
    marginTop: 40,
  },
  title: {
    fontSize: 40,
    fontWeight: "bold",
  },
  header: {
    fontSize: 48,
    color: "#FFFFFF",
    paddingBottom: 10,
    marginBottom: 40,
  },
  mainFormat: {
    marginTop: 30,
    marginBottom: 100,
  },

  textInput: {
    fontSize: 24,
    alignSelf: "stretch",
    height: 40,
    color: "#FFFFFF",
    borderBottomColor: "#fd9644",
    borderBottomWidth: 2,
    marginTop: 15,
  },
  button: {
    width: 325,
    alignSelf: "center",
    alignItems: "center",
    padding: 20,
    backgroundColor: "#fd9644",
    marginTop: 5,
  },
  buttonText: {
    fontSize: 20,
    color: "#FFFFFF",
    fontWeight: "bold",
  },
});
