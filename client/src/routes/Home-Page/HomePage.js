import React from 'react';
import { Typography } from '@mui/material';
import {Box, Grid} from '@mui/material';
import DonutChart from '../../components/DonutChart/DonutChart';
import { useOutletContext } from "react-router-dom";
const HomePage = () => {
  const [boxHeight, setBoxHeight] = useOutletContext();
  console.log(boxHeight);
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
    <Grid container sx={{ height: boxHeight }}>
      <Grid item xs={7} sx={{ height: '100%', border: '1px solid black' }}>
      
      <DonutChart data={data}/>
      </Grid>
      <Grid item xs={5} sx={{ height: '100%', border: '1px solid black' }}>
        <div>Second Grid item (30% width)</div>
      </Grid>
    </Grid>


    
  );
};

export default HomePage;