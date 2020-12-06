import { resolveSoa } from 'dns';
import express, {Application, Request, Response, NextFunction} from 'express';

// init express
const app: Application = express();

// setup routes 
app.get('/', (req: Request, res: Response, next: NextFunction) => {
    res.send('Zephyr-one ready to serve')
})

// Setup server with port 5000
app.listen(5000, () => console.log('Server running ...'))