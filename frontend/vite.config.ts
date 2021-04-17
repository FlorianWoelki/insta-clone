import { defineConfig } from 'vite';
import reactRefresh from '@vitejs/plugin-react-refresh';
// @ts-ignore
import reactSvgPlugin from 'vite-plugin-react-svg';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [reactRefresh(), reactSvgPlugin()],
});
