import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#4caf50',
      light: '#80e27e',
      dark: '#087f23',
    },
    background: {
      default: '#121212',
      paper: '#2d2d2d',
    },
    text: {
      primary: '#ffffff',
      secondary: '#b0b0b0',
      disabled: '#757575',
    },
    divider: '#424242',
  },
  typography: {
    fontFamily: 'inherit', // Inherit font from global styles
  },
});

export default theme;
