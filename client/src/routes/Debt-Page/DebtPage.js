import { Grid, Box, FormControl, Stack, InputLabel, FilledInput, InputAdornment, TextField, MenuItem, Typography, Button } from "@mui/material";
import { useState } from "react";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
// import MyCalendar from "../../components/Calendar/Calendar";
import "react-calendar/dist/Calendar.css";
import Calendar from 'react-calendar';
import '../../components/Calendar/Calendar.css'
import { useLoaderData, useNavigate } from "react-router";
import { sendDebt } from "../../api/Debt/SendDebt";
import DebtTable from "../../components/Debt-Table/DebtTable";

const mergeAndSortDebtsLoansByDate = (debtArray, loanArray) => {
  const debtData = debtArray.length > 0 ? debtArray.map((item) => ({
    ID: 'debt_'+item.id,
    date: item.date,
    description: item.description,
    amount: item.amount,
    type_id: item.type_id,
    type: "Debt",
    status: item.status,
  })) : [];

  const loanData = loanArray.length > 0 ? loanArray.map((item) => ({
    ID: 'loan_' + item.id,
    date: item.date,
    description: item.description,
    amount: item.amount,
    type_id: item.type_id,
    type: "Loan",
    status: item.status,
  })) : [];

  const mergedArray = [...debtData, ...loanData];

  return mergedArray.sort((a, b) => new Date(b.date) - new Date(a.date));
};

const DebtPage = () => {
    const navigate = useNavigate();
    const {debtData, debtTypesData} = useLoaderData();
    const [date, setDate] = useState(new Date());
    const [debtLoan, setDebtLoan] = useState('');
    const [debtType, setDebtType] = useState('');
    const [amount, setAmount] = useState();
    const [debtDescription, setDebtDescription] = useState('');
    const mergedData = mergeAndSortDebtsLoansByDate(debtData.DebtArr, debtData.LoanArr);
    const query = new URLSearchParams(location.search);

  //FILTERS
  const handleButtonClick = (e, value) => {
    // e.preventDefault();
    if (value) {
      // Update the URL with the new page number
      const newQuery = new URLSearchParams(query);
      newQuery.set('filter', value);
      navigate({ search: newQuery.toString() }); // navigate to update state values with useEffect
    } 
  };    
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
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'space-evenly'
          }}
        >
          
        <Box
          sx={{
            height: '5%',
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
          {['Year', 'Month', 'Week', 'Day'].map((period) => (
            <Button
              key={period}
              variant="contained"
              color="secondary"
              onClick={(e) => handleButtonClick(e, period.toLowerCase())}
              sx={{ border: '1px solid black', fontWeight: '600', color: '#0C1017', width: '100px' }}
            >
              {period}
            </Button>
          ))}
        </Stack>

      </Box>
          
          {/* DEBT DATA BOX */}
          <Box
            sx={{
              height: '90%',
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
            <DebtTable mergedData={mergedData}/>
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
                height: '35%',
                width: '80%',
                
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'space-evenly',
              }}
            >
            {/* CALENDAR */}
            <Calendar onChange={setDate} value={date} />

          </Box>
          <Box
              sx={{
                height: '60%',
                width: '80%',
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
               <Typography variant="h5"> 
               {/* {date.toDateString()} */}
               Қарыздар
               </Typography>
               <TextField 
                    id="filled-basic" 
                    label="Description"
                   
                    variant="filled"
                    onChange={(e)=> setDebtDescription(e.target.value)}
                    sx={{
                        width: '80%'
                    }}
                />
                {/* <TextField 
                    id="filled-basic" 
                    label="Amount" 
                    variant="filled"
                    onChange={(e)=> setAmount(e.target.value)}
                    sx={{
                        width: '80%'
                    }}
                /> */}

                <FormControl sx={{ width: '80%' }} variant="filled">
                    <InputLabel  required={true} htmlFor="filled-adornment-amount">Amount</InputLabel>
                    <FilledInput
                        id="filled-adornment-amount"
                        onChange={(e)=>setAmount(e.target.value)}
                        startAdornment={<InputAdornment position="start">₸</InputAdornment>}
                    />
                </FormControl>


                <TextField
                  id="outlined-select-spending"
                  variant="filled"
                  select
                  onChange={(e) => setDebtLoan(e.target.value)}
                  label="Select"
                  value={debtLoan}
                  sx={{ width: '80%' }}
                >
                  <MenuItem value="Debt">Debt</MenuItem>
                  <MenuItem value="Loan">Loan</MenuItem>
                </TextField>

                <TextField
                    id="outlined-select-spending"
                    variant="filled"
                    select
                    onChange={(e)=>setDebtType(e.target.value)}
                    value={debtType}
                    label="Select"
                    sx={{width: '80%'}}
                >
                    {debtTypesData && (debtTypesData.map((option) => (
                        // console.log(option.SpendingType)
                        <MenuItem key={option.id} value={option.id}>
                            {option.type}
                        </MenuItem>
                        ))
                    )}
                </TextField>

                <AddCircleOutlineIcon 
                    onClick={()=>sendDebt(date, debtLoan, debtType, amount, debtDescription, navigate)}
                    sx={{
                        fontSize:'50px',
                        '&:hover':{
                            cursor: 'pointer',
                            color: 'white'
                        },
                    }}
                />

          </Box>
        </Grid>
      </Grid>    )
}
export default DebtPage;