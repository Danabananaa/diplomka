import { Grid, Box, FormControl, InputLabel, FilledInput, InputAdornment, TextField, MenuItem, Typography, Button } from "@mui/material";
import { useState } from "react";
// import MyCalendar from "../../components/Calendar/Calendar";
import Calendar from 'react-calendar';
import '../../components/Calendar/Calendar.css'
import "react-calendar/dist/Calendar.css";

const DebtPage = () => {
    const [date, setDate] = useState(new Date());
    return (
        
        <Grid 
            container 
            alignItems="center"
            sx={{ 
            height: '100%'
            }}
        >
      {/* DEBT MAIN CONTAINER */}
        <Grid item xs={7} 
          sx={{
            height: '100%',
            // border:'1px solid black',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
          }}
        >
          {/* DEBT DATA BOX */}
          <Box
            sx={{
              height: '95%',
              width: '95%',
              backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
              boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
              borderRadius: '16px',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              border: '1px solid rgba(0, 0, 0, 0.25)',
              position: 'relative'
            }}
          >
            {/* DEBT DATA */}
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
          {/* CALENDAR BOX */}
          <Box
              sx={{
                height: '45%',
                width: '95%',
                backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
                boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
                borderRadius: '16px',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'space-evenly',
                border: '1px solid rgba(0, 0, 0, 0.25)'
              }}
            >
            {/* CALENDAR */}
            <Calendar onChange={setDate} value={date} />

          </Box>
          <Box
              sx={{
                height: '45%',
                width: '95%',
                backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
                boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
                borderRadius: '16px',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'space-evenly',
                border: '1px solid rgba(0, 0, 0, 0.25)'
              }}
            >
                {/* TODO Finish this part */}
               <Typography> 
               {date.toDateString()}
               </Typography>

          </Box>
        </Grid>
      </Grid>    )
}
export default DebtPage;