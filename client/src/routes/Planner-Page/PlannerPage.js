import { Grid, Box, Typography, FormControl, InputLabel, FilledInput, InputAdornment, TextField, MenuItem } from "@mui/material";
import { useLoaderData } from "react-router";
import PlannerTable from "../../components/Planner-Table/PlannerTable";
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import { useState } from "react";

const PlannerPage = () => {
    const {Types, Stats} = useLoaderData();
    const [spendingType, setSpendingType] = useState();
    const [spendingAmount, setSpendingAmount] = useState();


    return (
        
        <Grid 
            container 
            alignItems="center"
            sx={{ 
            height: '100%'
            }}
        >

      {/* <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">
          {"Деректер қатесі"}
        </DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
          Дұрыс емес деректер. Қайта толтырып көріңіз
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button variant="contained" onClick={handleClose} autoFocus sx={{border: "1px solid black"}}>
            Түсіндім
          </Button>
        </DialogActions>
      </Dialog> */}
      {/* BUDGET MAIN CONTAINER */}
        <Grid item xs={7} 
          sx={{
            height: '80%',
            // border:'1px solid black',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'space-evenly'
          }}
        >
          {/* BUDGET DATA BOX */}
          
          <Box
          sx={{
            height: '10%',
            width: '30%',
            backgroundImage: 'white',
            marginBottom: '40px',
            border: '2px solid black',
            boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
            borderRadius: '16px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            // border: '1px solid rgba(0, 0, 0, 0.25)',
            position: 'relative'
          }}
        >
            <Typography variant="h4">
                Planner
            </Typography>
      </Box>
          
          <Box
            sx={{
              height: '90%',
              width: '95%',
              // boxShadow: "0px 8px 10px rgba(0, 0, 0, 0.25)",
              // borderRadius: '16px',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              // border: '1px solid rgba(0, 0, 0, 0.25)',
              position: 'relative'
            }}
          >
              {/* TABLE */}
            {/* <BudgetTable mergedData={mergedData}/> */}
            <PlannerTable mergedData={Types}/>

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
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'space-evenly',
                border: '1px solid rgba(0, 0, 0, 0.25)'
              }}
            >
                {/* TOP HEADER */}
                <Typography variant="h4" >
                    Кірісті қосу
                </Typography>
              {/* TOP TYPE */}
                <TextField
                    id="outlined-select-currency"
                    variant="filled"
                    select
                    label="Таңдау"
                    onChange={(e)=>setSpendingType(e.target.value)}
                    required={true}
                    // value={incomeType}
                    
                    sx={{width: '80%'}}
                    >
                    {Types && Types.map((option) => (
                        <MenuItem key={option.ID} value={option.ID}>
                            {option.SpendingType}
                        </MenuItem>
                    ))}
                </TextField>
                
                {/* TOP AMOUNT */}
                <FormControl sx={{ width: '80%' }} variant="filled" required={true}>
                    <InputLabel htmlFor="filled-adornment-amount">Сумма</InputLabel>
                    <FilledInput
                        id="filled-adornment-amount"
                        onChange={(e)=>setSpendingAmount(e.target.value)}
                        inputProps={{ maxLength: 15 }}
                        startAdornment={<InputAdornment position="start">₸</InputAdornment>}
                    />
                </FormControl>
            {/* TOP BUTTON */}
                <AddCircleOutlineIcon 
                    // onClick={(e) =>{
                    //     if (/^\d+$/.test(income)) {
                    //       sendIncome(e, incomeType, income, incomeDescription, navigate)
                    //     } else{
                    //       handleClickOpen();
                    //     }  
                    //   }} 
                    sx={{
                        fontSize:'50px',
                        '&:hover':{
                            cursor: 'pointer',
                            color: 'white'
                        }
                    }}
                />

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
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'space-evenly',
                border: '1px solid rgba(0, 0, 0, 0.25)'
              }}
            >
              {/* BOTTOM HEADER */}
                <Typography variant="h4">
                    Шығысты қосу
                </Typography>
          </Box>
        </Grid>
      </Grid>

    )
}
export default PlannerPage;