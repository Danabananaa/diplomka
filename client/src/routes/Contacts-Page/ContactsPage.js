import { Grid, Box,Divider, FormControl, InputLabel, FilledInput, InputAdornment, TextField, MenuItem, Typography, Button } from "@mui/material";
import { Link } from "react-router-dom";
import contacts from '../../assets/images/contacts.png'
import git from '../../assets/images/git.png'
import facebook from '../../assets/images/facebook.png'
import insta from '../../assets/images/insta.png'
import twitter from '../../assets/images/twitter.png'
import linkedIn from '../../assets/images/linkedinn.png'
const ContactsPage = () => {
    
    return (
        <Grid container direction="row" sx={{ height: '100%' }}>
        {/* Full width grid item */}
        <Grid item xs={12} sx={{ height: '20%' }} display="flex" flexDirection="column" alignItems="center" justifyContent="center">
            <Typography variant="h2" textAlign="center">
                We'd Love yo Hear From You
            </Typography>
            <Typography variant="p" textAlign="center">
                Whether you are curious about our new features or upcoming features
            </Typography>
        </Grid>
  
        {/* 3 grid items side by side */}
    <Grid item container xs={12} sx={{ height: '80%' }} direction="row">
      <Grid item xs={4} 
          sx={{
            height: '100%',
            // border:'1px solid black',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
          }}
          >
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
               asd
            </Box>
        </Grid>
        <Grid item xs={4} 
        sx={{
            height: '100%',
            // border:'1px solid black',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
        }}
        >
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
            justifyContent: 'space-between',
            border: '1px solid rgba(0, 0, 0, 0.25)',
            position: 'relative',
            padding: '1rem',
          }}
        >
                    
            <Box
                sx={{
                display: 'flex',
                alignItems: 'center',
                }}
            >
                <img
                src={git}
                alt="My Image"
                style={{
                    width: '60px',
                    height: '60px',
                    marginRight: '8px',
                }}
                />
                <Typography variant="h6"><Link to="https://github.com/Danabananaa">https://github.com/Danabananaa </Link></Typography>
            </Box>

            <img
                src={contacts}
                alt="image"
                style={{
                    position: "absolute", // Move position property here
                    bottom: '50%',
                    left: '50%',
                    maxWidth: '100%',
                    maxHeight: '100%',
                    opacity: '0.6',
                    transform: 'translate(-50%, 50%)' // Add this line to center the image
                }}
            />   
            <Typography variant="h6">Bottom Typography</Typography>
        </Box>
        </Grid>
          <Grid item xs={4}
            sx={{
                height: '100%',
                // border:'1px solid black',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center'
            }}
          >
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
            <Typography variant="h4">Contacts</Typography>
            
            <Box
                sx={{
                display: 'flex',
                alignItems: 'center',
                }}
            >
                <img
                src={git}
                alt="My Image"
                style={{
                    width: '60px',
                    height: '60px',
                    marginRight: '8px',
                }}
                />
                <Typography variant="h6"><Link to="https://github.com/Danabananaa">https://github.com/Danabananaa </Link></Typography>
            </Box>
            <Box
                sx={{
                display: 'flex',
                alignItems: 'center',
                }}
            >
                <img
                src={linkedIn}
                alt="My Image"
                style={{
                    width: '60px',
                    height: '60px',
                    marginRight: '8px',
                }}
                />
                <Typography variant="h6"><Link to="https://github.com/Danabananaa">https://github.com/Danabananaa </Link></Typography>
            </Box>
            <Box
                sx={{
                display: 'flex',
                alignItems: 'center',
                }}
            >
                <img
                src={facebook}
                alt="My Image"
                style={{
                    width: '60px',
                    height: '60px',
                    marginRight: '8px',
                }}
                />
                <Typography variant="h6"><Link to="https://github.com/Danabananaa">https://github.com/Danabananaa </Link></Typography>
            </Box>
            
            <Box
                sx={{
                display: 'flex',
                alignItems: 'center',
                }}
            >
                <img
                src={twitter}
                alt="My Image"
                style={{
                    width: '60px',
                    height: '60px',
                    marginRight: '8px',
                }}
                />
                <Typography variant="h6"><Link to="https://github.com/Danabananaa">https://github.com/Danabananaa </Link></Typography>
            </Box>
            <Box
                sx={{
                display: 'flex',
                alignItems: 'center',
                }}
            >
                <img
                src={insta}
                alt="My Image"
                style={{
                    width: '60px',
                    height: '60px',
                    marginRight: '8px',
                }}
                />
                <Typography variant="h6"><Link to="https://github.com/Danabananaa">https://github.com/Danabananaa </Link></Typography>
            </Box>
        </Box>
        </Grid>
        </Grid>
      </Grid>
          )
}
export default ContactsPage;