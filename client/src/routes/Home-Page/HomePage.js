import React from 'react';
import { Paper, Typography } from '@mui/material';
import {Box, Grid} from '@mui/material';
import DonutChart from '../../components/DonutChart/DonutChart';
import {Button} from '@mui/material';
import {Stack} from '@mui/system';
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
          height: '90%',
          // border:'1px solid black',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'space-around'
        }}
      >

        <Box
          sx={{
            height: '10%',
            width: '95%',
            // backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
            // boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
            borderRadius: '16px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            // border: '1px solid rgba(0, 0, 0, 0.25)',
            position: 'relative'
          }}
        >
          <Typography variant="h4">Overall Statistics</Typography>

      </Box>
      <Box
          sx={{
            height: '10%',
            width: '95%',
            // backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
            // boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
            borderRadius: '16px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            // border: '1px solid rgba(0, 0, 0, 0.25)',
            position: 'relative'
          }}
        >
        <Stack direction="row" spacing={5}>
          <Button
            variant="contained"
            color="secondary"
            sx={{ border: "1px solid black", fontWeight: "600", color: "#0C1017", width: "100px" }}
          >
            Year
          </Button>
          <Button
            variant="contained"
            color="secondary"
            sx={{ border: "1px solid black", fontWeight: "600", color: "#0C1017", width: "100px" }}
          >
            Month
          </Button>
          <Button
            variant="contained"
            color="secondary"
            sx={{ border: "1px solid black", fontWeight: "600", color: "#0C1017", width: "100px" }}
          >
            Week
          </Button>
          <Button
            variant="contained"
            color="secondary"
            sx={{ border: "1px solid black", fontWeight: "600", color: "#0C1017", width: "100px" }}
          >
            Day
          </Button>
        </Stack>

      </Box>
        {/* PIE BOX */}
        <Box
          sx={{
            height: '80%',
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
          {/* PIE */}
          <svg
            viewBox="0 0 256 256"
            xmlns="http://www.w3.org/2000/svg"
            style={{
              position: 'absolute', // set position to absolute for precise placement
              top: '50%', // move to the middle of the container
              left: '50%',
              height: '30%',
              width: '30%',
              transform: 'translate(-50%, -50%)', // center the SVG
            }}
          >     
            <g fill="#000" strokeWidth="1">
              <path d="M127.857 1.407c-69.725 0-126.45 56.725-126.45 126.45 0 69.724 56.725 126.45 126.45 126.45 69.724 0 126.45-56.726 126.45-126.45 0-69.725-56.726-126.45-126.45-126.45zm0 236.04c-60.43 0-109.59-49.161-109.59-109.59s49.16-109.59 109.59-109.59 109.59 49.16 109.59 109.59-49.161 109.59-109.59 109.59z" />
              <path d="M184.357 94.738H71.356c-4.656 0-8.43 3.774-8.43 8.43s3.774 8.43 8.43 8.43h48.07v87.914a8.43 8.43 0 008.43 8.43 8.43 8.43 0 008.43-8.43v-87.914h48.071c4.656 0 8.43-3.774 8.43-8.43s-3.77-8.43-8.43-8.43zM71.356 81.492H184.36c4.656 0 8.43-3.774 8.43-8.43s-3.774-8.43-8.43-8.43H71.356c-4.656 0-8.43 3.773-8.43 8.43s3.774 8.43 8.43 8.43z" />
            </g>
        </svg>
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
              height: '45%',
              width: '80%',
              backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
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
              height: '45%',
              width: '80%',
              backgroundImage: 'linear-gradient(0deg, #c2b6df 10%, #cdb2bd 90%)',
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