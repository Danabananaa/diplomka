import React from 'react';
import { List, ListItem, ListItemIcon, ListItemText, Drawer, Toolbar, Box } from '@mui/material';
import { Dashboard as DashboardIcon, BarChart as BarChartIcon, AttachMoney as AttachMoneyIcon, EventNote as EventNoteIcon, CreditCard as CreditCardIcon, Settings as SettingsIcon, Info as InfoIcon, Mail as MailIcon } from '@mui/icons-material';
import {Divider} from '@mui/material';
import {ListItemButton} from '@mui/material';

export const drawer = (
    <div>
      <Toolbar />
      <Divider />
      <List>
        {['Statistics', 'Budget', 'Planner', 'Debt'].map((text, index) => (
          <ListItem key={text} disablePadding>
            <ListItemButton>
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
            <ListItemButton>
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