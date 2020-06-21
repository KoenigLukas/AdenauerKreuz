import React from "react";
import {
  StyleSheet,
  Text,
  View,
  TextInput,
  TouchableOpacity,
} from "react-native";

export default function Register() {
  return (
    <View style={styles.register}>
      <Text style={styles.title}>New User?</Text>
      <Text style={styles.subTitle}>Register Here:</Text>
      <View style={styles.mainFormat}>
        <TextInput
          style={styles.textInput}
          placeholder="E-Mail:"
          underlineColorAndroid={"transparent"}
        />
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
        <TextInput
          style={styles.textInput}
          placeholder="Enter Password again"
          secureTextEntry={true}
          underlineColorAndroid={"transparent"}
        />
      </View>

      <TouchableOpacity style={styles.button}>
        <Text style={styles.buttonText}>Sign Up</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  register: {
    alignSelf: "stretch",
    marginTop: 0,
  },
  header: {
    fontSize: 48,
    color: "#000000",
    paddingBottom: 10,
    marginBottom: 40,
  },
  title: {
    fontSize: 40,
    fontWeight: "bold",
  },
  subTitle: {
    fontSize: 20,
    fontWeight: "bold",
  },
  mainFormat: {
    marginTop: 50,
    marginBottom: 75,
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
