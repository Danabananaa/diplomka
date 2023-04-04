import React from 'react';
import { Paper, Typography } from '@mui/material';
import {Box, Grid} from '@mui/material';
import DonutChart from '../../components/DonutChart/DonutChart';
import { display } from '@mui/system';
// import { useOutletContext } from "react-router-dom";
const HomePage = () => {
  // const [mainHeight, setMainHeight] = useOutletContext();
  // console.log(mainHeight);
  const data = [
    {
      "id": "scala",
      "label": "scala",
      "value": 102,
      "color": "hsl(1, 70%, 50%)"
    },
    {
      "id": "haskell",
      "label": "haskell",
      "value": 265,
      "color": "hsl(242, 70%, 50%)"
    },
    {
      "id": "php",
      "label": "php",
      "value": 594,
      "color": "hsl(228, 70%, 50%)"
    },
    {
      "id": "ruby",
      "label": "ruby",
      "value": 156,
      "color": "hsl(37, 70%, 50%)"
    },
    {
      "id": "c",
      "label": "c",
      "value": 546,
      "color": "hsl(232, 70%, 50%)"
    }
  ]
  return (
    <Grid 
      container 
      alignItems="center"
      sx={{ 
        height: '100%'
      }}
    >
    {/* PIE CONTAINER */}
      <Grid item xs={7} 
        sx={{
          height: '80%',
          // border:'1px solid black',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center'
        }}
      >
        {/* PIE BOX */}
        <Box
          sx={{
            height: '95%',
            width: '95%',
            backgroundImage: 'linear-gradient(0deg, #cdb2bd 10%, #c2b6df 90%)',
            boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
            borderRadius: '16px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            border: '1px solid rgba(0, 0, 0, 0.25)'
          }}
        >
          {/* PIE */}
          <DonutChart data={data}/>
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
              height: '40%',
              width: '80%',
              backgroundImage: 'linear-gradient(0deg, #cdb2bd 10%, #c2b6df 90%)',
              boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
              borderRadius: '16px',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              border: '1px solid rgba(0, 0, 0, 0.25)'
            }}
          >
            INCOME
        </Box>
        {/* BOTTOM BOX */}
        <Box
            sx={{
              height: '40%',
              width: '80%',
              backgroundImage: 'linear-gradient(0deg, #cdb2bd 10%, #c2b6df 90%)',
              boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
              borderRadius: '16px',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              border: '1px solid rgba(0, 0, 0, 0.25)'
            }}
          >
            SHIGISTAR
        </Box>
      </Grid>
    </Grid>
    
  );
};

export default HomePage;