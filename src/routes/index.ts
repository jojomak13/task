import { Router } from 'express';

const router = Router();

router.get('/', async (_req, res) => {
    res.json({ msg: 'welcome' });
});

export default router;