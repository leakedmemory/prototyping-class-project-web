import { router } from "expo-router";
import { Text, View, StyleSheet } from "react-native";
import { useSession } from "../contexts/authentication";
import Input from "../components/input";
import FontAwesome from "@expo/vector-icons/FontAwesome";
import Button from "../components/button";

export default function SignIn() {
  const { signIn } = useSession();
  return (
    <View style={styles.root}>
      <View style={styles.loginBox}>
        <Text style={styles.loginText}>Login</Text>
        <View style={styles.boxInputs}>
          <Input
            icon={<FontAwesome name="user" size={24} color="#828282" />}
            style={{ marginBottom: 10 }}
          />
          <Input icon={<FontAwesome name="lock" size={24} color="#828282" />} />
        </View>
        <Button
          onPress={() => {
            signIn();
            router.replace("/");
          }}
          text={"Entrar"}
        />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  root: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "#ff7260",
  },
  loginBox: {
    backgroundColor: "white",
    alignItems: "center",
    width: "70%",
    borderRadius: 10,
    padding: 30,
    justifyContent: "space-between",
  },
  loginText: {
    fontSize: 48,
    fontWeight: "800",
    color: "#353535",
    marginBottom: 10,
  },
  boxInputs: {
    width: "100%",
    marginBottom: 10,
  },
});
