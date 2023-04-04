import { redirect } from "react-router";


export const homeLoader = () => {
    
    const token = localStorage.getItem('token');

    if (token) {
        return null
    } else {
        return (redirect("/signin"))
    }
}


export const typesData =( async () => {
        try {
        //   const url = new URL(request.url);
        //   const searchParams = url.searchParams;
        //   const id = searchParams.get('id');
            
          const [incomeResponse, spendingResponse] = await Promise.all([
            fetch(`/income/type`, {
              headers: {
                Accept: "application/json",    
            },
              method: "GET",
            }),
            fetch(`/spending/type`, {
              headers: {
                Accept: "application/json",
              },
              method: "GET",
            })
          ]);
            
          if (!incomeResponse.ok) {
            const error = new Error(`Could not fetch the post. Status: ${incomeResponse.statusText}`);
            error.status = incomeResponse.status;
            throw error;
          }
          
          if (!spendingResponse.ok) {
            const error = new Error(`Could not fetch the comments. Status: ${spendingResponse.statusText}`);
            error.status = spendingResponse.status;
            throw error;
          }
          
          const incomeTypes = await incomeResponse.json();
          const spendingTypes = await spendingResponse.json();
          return { incomeTypes, spendingTypes };
        } catch (error) {
          console.log(error);
          throw error;
        }
       
})  