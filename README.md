# pack-picker

### Introduction
This project includes an API which calculates the number of packs to ship to a customer based on a given list of set pack sizes. This repo includes the backend work for the pack calculation, the front end can be found [here](https://github.com/rachelcurran/pack-picker-front.git).

___

### API Endpoints
| HTTP Verbs | Endpoints | Action | Notes 
| --- | --- | --- | --- |
| POST | /packs | To calculate the number of packs to fulfil an order | |
| GET | /packs-old | To calculate the number of packs to fulfil an order | This endpoint doesn't always provide as accurate results and shouldn't be used |
| POST | /pack-sizes | Sets the pack sizes to be used with `/packs-old` | This endpoint is only used by `/packs-old` and is no longer required |
| GET | /pack-sizes | Gets the pack sizes to be used with `/packs-old` | This endpoint is only used by `/packs-old` and is no longer required |


The application should be available at `https://whispering-hollows-05368-f071ed37f19c.herokuapp.com`. 

Please let me know if you have any issues accessing this.

Example request: 

`POST https://whispering-hollows-05368-f071ed37f19c.herokuapp.com/packs`

    {
      "numberOfItems": 100,
      "packSizes" : [250, 500, 1000, 2000, 5000]
    }

___

### Notes
* Unit tests for all calculations are included.
* The front end may include a few small bugs so please do any major testing of the calculations at the API level.
* Due to the method of calculation, certain pack sizes and number of items in order can take a very long time to process (and most likely will time out). This will likely occur when using a large number for items in order along with small pack sizes 


    
