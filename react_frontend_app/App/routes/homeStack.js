import { createStackNavigator } from "react-navigation-stack";
import { createAppContainer } from "react-navigation";
import Registration from "../components/Register";
import Login from "../components/Login";
import Mainmenu from "../components/Mainmenu";
import AdenauerKreuz from "../components/AdenauerKreuz";

const screens = {
  AdenauerKreuz: {
    screen: AdenauerKreuz,
  },
  Login: {
    screen: Login,
  },
  Registration: {
    screen: Registration,
  },
  Mainmenu: {
    screen: Mainmenu,
  },
  AdenauerKreuz: {
    screen: AdenauerKreuz,
  },
};

const HomeStack = createStackNavigator(screens);

export default createAppContainer(HomeStack);
