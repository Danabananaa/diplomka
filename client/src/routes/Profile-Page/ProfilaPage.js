import { Grid, Box, FormControl, List, ListItem, InputLabel, FilledInput, InputAdornment, TextField, MenuItem, Typography, Button, Divider, Avatar } from "@mui/material";
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import { useLoaderData } from "react-router";
import { sendIncome } from "../../api/Budget/SendIncome";
import { useState } from "react";
import { useNavigate } from "react-router";
import { sendSpending } from "../../api/Budget/SendSpending";
import girlImage from '../../assets/images/girl.png'
import bulb from '../../assets/images/bulb.png'
import ImageUploadForm from "../../components/Image-Upload-Form/ImageUploadForm";
import { fetchImage } from "../../api/Avatar/Avatar";
import { useEffect } from "react";
import { useSelector } from "react-redux";

const ProfilePage = () => {
    const imageURL = useSelector((state) => state.image.imageURL);
    
    return (
        <Grid 
            container 
            alignItems="center"
            sx={{ 
            height: '100%'
            }}
        >
      {/* BUDGET MAIN CONTAINER */}
        <Grid item xs={7} 
          sx={{
            height: '100%',
            // border:'1px solid black',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
          }}
        >
          {/* BUDGET DATA BOX */}
          <Box
            sx={{
                height: '95%',
                width: '95%',
                backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
                boxShadow: '0px 8px 10px rgba(0, 0, 0, 0.25)',
                borderRadius: '16px',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'space-evenly',
                border: '1px solid rgba(0, 0, 0, 0.25)',
                position: 'relative',
                padding: '1rem',
            }}
            >
            <Typography variant="h4" component="h2">
                Profile Page
            </Typography>

            {imageURL && (
            <>
                <Avatar src={imageURL} />
            </>
            )}
            <ImageUploadForm/>
            
        </Box>
        </Grid>
        {/* RIGHT CONTAINER */}
        <Grid item xs={5} 
          sx={{ 
            height: '100%',
            // border: '1px solid black',
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'space-evenly',
            alignItems: 'center'
          }}
        >
          {/* TOP BOX */}
          <Box
            sx={{
                height: '45%',
                width: '80%',
                backgroundColor: 'White',
                boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
                borderRadius: '16px',
                display: 'flex',
                flexDirection: 'row',
                alignItems: 'flex-end',
                border: '1px solid rgba(0, 0, 0, 0.25)',
                position: 'relative',
            }}
            >
            <img
                src={girlImage}
                alt="image"
                style={{
                marginBottom: 0, // Add this property
                marginLeft: 0, 
                maxWidth: '80%',
                maxHeight: '80%',
                objectFit: 'contain',
                }}
            />
            
            
        </Box>
          {/* BOTTOM BOX */}
          <Box
              sx={{
                height: '45%',
                width: '80%',
                boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
                borderRadius: '16px',
                display: 'flex',
                flexDirection: 'row',
                alignItems: 'flex-end',
                border: '1px solid rgba(0, 0, 0, 0.25)',
                position: 'relative',

              }}
            >
              {/* BOTTOM INFO */}
                <Box sx={{pl:2,alignSelf:'center', flexGrow: "0.9"}}>

                    <Typography variant="h5" mb={3} mt={3} textAlign="center">
                        Used technologies
                    </Typography>
                    <Divider/>
                    <List sx={{ listStyleType: 'disc', pl: 4 , fontSize:'1.2rem'}}>
                        <ListItem sx={{ display: 'list-item' }}>Nivo Charts</ListItem>
                        <ListItem sx={{ display: 'list-item' }}>React Calendar</ListItem>
                        <ListItem sx={{ display: 'list-item' }}>SQLite</ListItem>
                        <ListItem sx={{ display: 'list-item' }}>Banana</ListItem>
                    </List>
                </Box>
              <img
                src={bulb}
                alt="image"
                style={{
                marginBottom: 0, // Add this property
                marginRight: 0, 
                maxWidth: '70%',
                maxHeight: '70%',
                objectFit: 'contain',
                }}
            />
        
          </Box>
        </Grid>
      </Grid>
      
    )
}
export default ProfilePage;