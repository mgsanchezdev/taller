import { View, Text, StyleSheet } from 'react-native';

export default function Login() {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Login</Text>
      <Text style={styles.subtitle}>Pantalla de inicio de sesión</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, justifyContent: 'center', alignItems: 'center', padding: 24 },
  title: { fontSize: 22, fontWeight: 'bold', marginBottom: 8 },
  subtitle: { color: '#666' },
});
