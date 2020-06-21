import { createStackNavigator } from "react-navigation-stack";
import { createAppContainer } from "react-navigation";
import Registration from "../components/Register";
import Login from "../components/Login";

const screens = {
  Login: {
    screen: Login,
  },
  Registration: {
    screen: Registration,
  },
};

const HomeStack = createStackNavigator(screens);

export default createAppContainer(HomeStack);
