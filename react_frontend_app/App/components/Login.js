import React from "react";
import {
  StyleSheet,
  Text,
  View,
  TextInput,
  TouchableOpacity,
  KeyboardAvoidingView,
  AsyncStorage,
} from "react-native";

export default class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
      token: "",
    };
  }

  componentDidMount() {
    this._loadInitialState().done();
  }

  _loadInitialState = async () => {
    var value = await AsyncStorage.getItem("user");
    if (value !== null) {
      this.props.navigation.navigate("Profile");
    }
  };

  render() {
    return (
      <View style={styles.logIn}>
        <KeyboardAvoidingView behavior="padding">
          <Text style={styles.title}>Hello.</Text>
          <Text style={styles.title}>Welcome Back</Text>

          <View style={styles.mainFormat}>
            <TextInput
              style={styles.textInput}
              placeholder="User-Name"
              underlineColorAndroid={"transparent"}
              onChangeText={(username) => this.setState({ username })}
            />
            <TextInput
              style={styles.textInput}
              placeholder="Enter Password"
              secureTextEntry={true}
              underlineColorAndroid={"transparent"}
              onChangeText={(password) => this.setState({ password })}
            />
          </View>
          <TouchableOpacity style={styles.button} onPress={this.login}>
            <Text style={styles.buttonText}>Log In</Text>
          </TouchableOpacity>

          <TouchableOpacity
            onPress={this.registerButtonHandler}
            style={styles.button}
          >
            <Text style={styles.buttonText}>Register</Text>
          </TouchableOpacity>
        </KeyboardAvoidingView>
      </View>
    );
  }

  registerButtonHandler = () => {
    this.props.navigation.navigate("Registration");
  };

  login = () => {
    if (this.state.username === "" || this.state.password === "") {
      alert("Username / Password is Incorrect");
    } else {
      this.props.navigation.navigate("Mainmenu", {
        username: this.state.username,
      });
    }

    /*
    fetch(" http://localhost:8000/api/v1/login", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: this.state.username,
        password: this.state.password,
      }),
    })
      .then((response) => response.json())
      .then((res) => {
        if (res.success === true) {
          AsyncStorage.setItem("user", res.user);
          this.props.navigation.navigate("");
        } else {
          alert(res.message);
        }
      })
      .done();
      */
  };
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
    color: "#000",
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
