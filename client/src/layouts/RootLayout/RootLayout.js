import { Outlet, useLoaderData, useLocation } from 'react-router-dom';
import { Box, CssBaseline, AppBar, Toolbar, Typography, IconButton, Drawer, Avatar } from '@mui/material';
import { Menu as MenuIcon, Dashboard as DashboardIcon } from '@mui/icons-material';
import RootDrawer from '../../components/Root-Drawer/RootDrawer';
import { MainDrawer } from '../../components/Root-Drawer/RootDrawer';
import { useState } from 'react';
import { Search } from '@mui/icons-material';
import {InputBase} from '@mui/material';
import { useSelector } from 'react-redux';

const drawerWidth = 240;

const RootLayout = () => {


  const username = useSelector((state) => state.auth.username)
  const location = useLocation();
  const path = location.pathname.split('/')[1];
  return (
    <Box sx={{ display: 'flex', minHeight: '100vh', backgroundColor:"primary.light"}}>
      <CssBaseline />
      {/* APPBAR */}
      <AppBar position="fixed" sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px`, backgroundColor: 'primary.lighter', color: 'black'}}>
        <Toolbar>
          <Box sx={{ display: 'flex', alignItems: 'center', flexGrow: 1 }}>
            <Box sx={{ marginLeft: 2, display: 'flex', alignItems: 'center' }}>
              <InputBase
                placeholder="Search…"
                inputProps={{ 'aria-label': 'search' }}
                sx={{ marginLeft: 1, backgroundColor: "primary.light", borderRadius:4, p:0.7 }}
              />
              <Search/>
            </Box>
          </Box>
          <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <Typography variant="body1" component="span" sx={{ marginRight: 1 }}>
              {username ? username : 'User'}
            </Typography>
            <Avatar alt="User Avatar" />
          </Box>
        </Toolbar>
      </AppBar>
        {/* DRAWER */}
      <Box
        component="nav"
        sx={{ width: { sm: drawerWidth }, flexShrink: { sm: 0 }}}
        aria-label="Navigation-Bar"
      >
        <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
            backgroundImage: "linear-gradient(0deg, #cdb2bd 10%, #c2b6df 90%)",
            border: "none"
          },
        
        }}
        variant="permanent"
        anchor="left"
      >
          <MainDrawer path={path}/>
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