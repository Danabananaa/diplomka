import { useState } from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import AccountCircle from '@mui/icons-material/AccountCircle';
import Switch from '@mui/material/Switch';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormGroup from '@mui/material/FormGroup';
import MenuItem from '@mui/material/MenuItem';
import Menu from '@mui/material/Menu';
import { Avatar } from '@mui/material';
import { useNavigate } from 'react-router';
import { useLoaderData } from 'react-router';
import { useEffect } from 'react';
import { fetchImage } from '../../api/Avatar/Avatar';
import { useDispatch } from 'react-redux';
import { fetchImageSuccess } from '../../utils/reducers/auth';
import { fetchAndSetImageURL } from '../../utils/reducers/avatar';
import { useSelector } from 'react-redux';

export default function MenuAppBar({drawerWidth}) {
  const {avatar} = useLoaderData(); // Avatar from the RootLayout loader
  const imageURL = useSelector((state) => state.image.imageURL);

  const dispatch = useDispatch();

    useEffect(() => {
        if (avatar) {
          dispatch(fetchAndSetImageURL(avatar.name));
        }
      }, [avatar, dispatch]);

  
  const [auth, setAuth] = useState(true);
  const [anchorEl, setAnchorEl] = useState(null);
  const navigate = useNavigate();

  const handleChange = (event) => {
    setAuth(event.target.checked);
  };

  const handleMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleProfile = () => {
    setAnchorEl(null);
    navigate("/profile");
  }

  return (
   
      <AppBar
        
        sx={{
        width: `calc(100% - ${drawerWidth}px)`,
        ml: `${drawerWidth}px`,
        backgroundColor: 'primary.lighter',
        color: 'black'
      }}
      >
        <Toolbar>
              <Avatar src={imageURL} onClick={handleMenu}/>
              <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                <MenuItem onClick={handleProfile}>Profile</MenuItem>
                <MenuItem onClick={handleClose}>My account</MenuItem>
              </Menu>
           
        </Toolbar>
      </AppBar>
  );
}