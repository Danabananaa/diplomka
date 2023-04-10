import { List, ListItem, ListItemIcon, ListItemText, Drawer, Toolbar, Box, Typography, Button } from '@mui/material';
import { Dashboard as DashboardIcon, BarChart as BarChartIcon, AttachMoney as AttachMoneyIcon, EventNote as EventNoteIcon, CreditCard as CreditCardIcon, Settings as SettingsIcon, Info as InfoIcon, Mail as MailIcon } from '@mui/icons-material';
import {Divider} from '@mui/material';
import {ListItemButton} from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router';
import LogoutIcon from '@mui/icons-material/Logout';
import { signOutHandler } from '../../api/Authorization/Authorization';
import { useDispatch } from 'react-redux';
export const MainDrawer = ({path}) =>{
  
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const handleNav = (text) => {
    const newPath = text.replace(/\s/g, ''); // Removing spaces from a button text
    navigate(`/${newPath.toLowerCase()}`); // changing the url to the specified nav button text
  }
  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        height: '100%', // Make sure the Box component expands to the full height of its container
      }}
    >
      <Box>

        <Toolbar />
        <Divider />
        <List >
          {['Statistics', 'Budget', 'Planner', 'Debt'].map((text, index) => (
            <ListItem key={text} disablePadding>
                <ListItemButton
                  onClick={() => handleNav(text)}
                  sx={{
                    backgroundColor: (text.toLowerCase().replace(/\s/g, '') === path.toLowerCase().replace(/\s/g, '')) && "primary.light",
                    border: "1px solid transparent",
                    borderColor: (text.toLowerCase().replace(/\s/g, '') === path.toLowerCase().replace(/\s/g, '')) ? "black transparent black black" : "none",
                    // my: 1,
                    py: 2,
                    borderTopLeftRadius: 20,
                    borderBottomLeftRadius: 20,
                    '&:hover': {
                      backgroundColor: 'primary.main',
                    },
                  }}
                > 
                  <ListItemIcon>
                  {index % 2 === 0 ? <InfoIcon /> : <MailIcon />}
                </ListItemIcon>
                <ListItemText primary={text} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
        <Divider />
        <List>
          {['Settings', 'About Us', 'Contacts'].map((text, index) => (
            <ListItem key={text} disablePadding>
              <ListItemButton
                onClick={() => handleNav(text)}
                sx={{
                  backgroundColor: (text.toLowerCase().replace(/\s/g, '') === path.toLowerCase().replace(/\s/g, '')) && "primary.light",
                  border: "1px solid transparent",
                  borderColor: (text.toLowerCase().replace(/\s/g, '') === path.toLowerCase().replace(/\s/g, '')) ? "black transparent black black" : "none",
                    // marginY: 1,
                  py: 2,
                  borderTopLeftRadius: 20,
                  borderBottomLeftRadius: 20,
                  '&:hover': {
                    backgroundColor: 'primary.main',
                  },
                }}
              >
                  <ListItemIcon>
                  {index % 2 === 0 ? <InfoIcon /> : <MailIcon />}
                </ListItemIcon>
                <ListItemText primary={text} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>

      </Box>
        
        <Box
        onClick={signOutHandler(dispatch, navigate)}
        sx={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          mb: 6,
          "&:hover": {
            color: "aliceblue",
            cursor: "pointer"
          },
          "&:active":{
            color: "#FF2020",
          }
        }}
        >
        
        <LogoutIcon fontSize='large'/><Typography>Sign Out</Typography>
        
        </Box>
      </Box>
    );
} 