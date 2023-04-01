import { Outlet } from 'react-router-dom';
import { Box, CssBaseline, AppBar, Toolbar, Typography, IconButton, Drawer } from '@mui/material';
import { Menu as MenuIcon, Dashboard as DashboardIcon } from '@mui/icons-material';
import RootDrawer from '../../components/Root-Drawer/RootDrawer';
import { drawer } from '../../components/Root-Drawer/RootDrawer';
import { useState } from 'react';

const drawerWidth = 240;

const RootLayout = (props) => {
  const [open, setOpen] = useState(false);
  const { window } = props;
  const [mobileOpen, setMobileOpen] = useState(false);
  const handleDrawerToggle = () => {
    setOpen(!open);
  };
  const container = window !== undefined ? () => window().document.body : undefined;
  return (
    <Box sx={{ display: 'flex', minHeight: '100vh'}}>
      <CssBaseline />
      {/* APPBAR */}
      <AppBar position="fixed" sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px`, backgroundColor: 'white', color: 'black'}}>
        <Toolbar>
          <IconButton edge="start" color="inherit" aria-label="menu" onClick={handleDrawerToggle}>
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" component="div">
            Dashboard
          </Typography>
        </Toolbar>
      </AppBar>
        {/* DRAWER */}
      <Box
        component="nav"
        sx={{ width: { sm: drawerWidth }, flexShrink: { sm: 0 } }}
        aria-label="mailbox folders"
      >
        <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
        }}
        variant="permanent"
        anchor="left"
      >
          {drawer}
        </Drawer>
      </Box>
      {/* MAIN */}
      <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
        <Toolbar />
        <Outlet />
      </Box>
    </Box>
  );
};

export default RootLayout;