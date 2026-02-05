import { Stack } from 'expo-router';

export default function RootLayout() {
  return (
    <Stack screenOptions={{ headerShown: false }}>
      <Stack.Screen name="index" />
      <Stack.Screen name="login" />
      <Stack.Screen name="reception" />
      <Stack.Screen name="inspection" />
      <Stack.Screen name="signature" />
      <Stack.Screen name="history" />
    </Stack>
  );
}
