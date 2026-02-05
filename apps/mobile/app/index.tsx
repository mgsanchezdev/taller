import { Link, useRouter } from 'expo-router';
import { View, Text, StyleSheet, Pressable } from 'react-native';

export default function Home() {
  const router = useRouter();
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Workshop Platform</Text>
      <Text style={styles.subtitle}>App móvil</Text>
      <Pressable style={styles.button} onPress={() => router.push('/login')}>
        <Text style={styles.buttonText}>Iniciar sesión</Text>
      </Pressable>
      <Link href="/reception" asChild>
        <Pressable style={styles.link}>
          <Text style={styles.linkText}>Recepción</Text>
        </Pressable>
      </Link>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, justifyContent: 'center', alignItems: 'center', padding: 24 },
  title: { fontSize: 24, fontWeight: 'bold', marginBottom: 8 },
  subtitle: { color: '#666', marginBottom: 24 },
  button: { backgroundColor: '#22c55e', paddingVertical: 12, paddingHorizontal: 24, borderRadius: 8, marginBottom: 12 },
  buttonText: { color: '#fff', fontWeight: '600' },
  link: { padding: 12 },
  linkText: { color: '#22c55e' },
});
