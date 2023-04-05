import { redirect } from "react-router";


export const homeLoader = () => {
    
    const token = localStorage.getItem('token');

    if (token) {
        return null
    } else {
        return (redirect("/signin"))
    }
}


export const budgetData =( async () => {
  const token = localStorage.getItem('token');
  
  try {
        //   const url = new URL(request.url);
        //   const searchParams = url.searchParams;
        //   const id = searchParams.get('id');
            
          const [budgetTypeResponse, budgetStatsResponse] = await Promise.all([
            fetch(`/budget/type`, {
              headers: {
                Accept: "application/json",
                Authorization: `${token}`,
              },
              method: "GET",
              
            }),
            fetch(`/budget/stats`, {
              headers: {
                Accept: "application/json",
                Authorization: `${token}`,
              },
              method: "GET",
              
            })
          ]);
            
          if (!budgetTypeResponse.ok) {
            // if (budgetTypeResponse.status===401){
            //   return redirect('/signin')
            // }
            const error = new Error(`Could not fetch the post. Status: ${budgetTypeResponse.statusText}`);
            error.status = budgetTypeResponse.status;
            throw error;
          }
          
          if (!budgetStatsResponse.ok) {
            // if (budgetStatsResponse.status===401){
            //   return redirect('/signin')

            // }
            const error = new Error(`Could not fetch the comments. Status: ${budgetStatsResponse.statusText}`);
            error.status = budgetStatsResponse.status;
            throw error;
          }
          
          const Types = await budgetTypeResponse.json();
          const Stats = await budgetStatsResponse.json();
          return { Types, Stats };
        } catch (error) {
          console.log(error);
          throw error;
        }
})  

export const debtData =( async () => {
  const token = localStorage.getItem('token');
  try {
        //   const url = new URL(request.url);
        //   const searchParams = url.searchParams;
        //   const id = searchParams.get('id');
            
          const [debtResponse, debtTypesResponse] = await Promise.all([
            fetch(`/debt`, {
              headers: {
                Accept: "application/json",
                Authorization: `${token}`,
              },
              method: "GET",
              
            }),
            fetch(`/assliatype`, {
              headers: {
                Accept: "application/json",
                Authorization: `${token}`,
              },
              method: "GET",
              
            }),
          ]);
            
          if (!debtResponse.ok) {
            // if (budgetTypeResponse.status===401){
            //   return redirect('/signin')
            // }
            const error = new Error(`Could not fetch the post. Status: ${budgetTypeResponse.statusText}`);
            error.status = budgetTypeResponse.status;
            throw error;
          }
          if (!debtTypesResponse.ok) {
            // if (budgetTypeResponse.status===401){
            //   return redirect('/signin')
            // }
            const error = new Error(`Could not fetch the post. Status: ${budgetTypeResponse.statusText}`);
            error.status = budgetTypeResponse.status;
            throw error;
          }
          
          const debtData = await debtResponse.json();
          const debtTypesData = await debtTypesResponse.json();
          console.log(debtTypesData);
          return { debtTypesData };
        } catch (error) {
          console.log(error);
          throw error;
        }
})  