import { List, ListItem, ListItemIcon, ListItemText, Drawer, Toolbar, Box } from '@mui/material';
import { Dashboard as DashboardIcon, BarChart as BarChartIcon, AttachMoney as AttachMoneyIcon, EventNote as EventNoteIcon, CreditCard as CreditCardIcon, Settings as SettingsIcon, Info as InfoIcon, Mail as MailIcon } from '@mui/icons-material';
import {Divider} from '@mui/material';
import {ListItemButton} from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router';
export const MainDrawer = ({path}) =>{
  
  const navigate = useNavigate();

  const handleNav = (text) => {
    const newPath = text.replace(/\s/g, ''); // Removing spaces from a button text
    navigate(`/${newPath.toLowerCase()}`); // changing the url to the specified nav button text
  }
  return (
      <div>
        <Toolbar />
        <Divider />
        <List >
          {['Statistics', 'Budget', 'Planner', 'Debt'].map((text, index) => (
            <ListItem key={text} disablePadding>
                <ListItemButton
                  onClick={() => handleNav(text)}
                  sx={{
                    backgroundColor: (text.toLowerCase().replace(/\s/g, '') === path.toLowerCase().replace(/\s/g, '')) && "primary.light",
                    my: 1,
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
                  marginY: 1,
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
      </div>
    );
} 