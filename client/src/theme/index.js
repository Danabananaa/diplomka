import { createTheme } from '@mui/material/styles';

const shadesOfWhite = {
  lighter: '#fafafa',
  light: '#f5f5f5',
  main: '#ffffff',
  dark: '#e0e0e0',
  darker: '#d6d6d6',
  contrastText: '#000000',
};

const themeOptions = {
  palette: {
    mode: 'light',
    primary: shadesOfWhite,
    secondary: {
      main: '#f50057',
    },
  },
  spacing: 8,
  components: {
    MuiAppBar: {
      variants: [
        {
          props: { color: 'inherit' },
          style: {
            backgroundColor: '#689f38',
            color: '#fff',
          },
        },
      ],
    },
  },
};

const theme = createTheme(themeOptions);
export default theme;