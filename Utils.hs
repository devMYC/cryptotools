-- take 3 $ map fst $ dropWhile (\(a,_) -> a < 1000) $ primElem 4969 [2,3,23]
-- first three generators greater than 1000
primElem p factorsOfPMinusOne =
    filter
        (\(_, rs) -> all (\n -> n /= 1) rs)
        [(a, [mod (a^((p-1) `div` f)) p | f <- factorsOfPMinusOne]) | a <- [1..p-1]]


-- 7^x \equiv 11 mod 13
-- x = dlp 7 11 13
dlp b t m =
    case take 1 $ dropWhile (\(_, r) -> r /= t) [(x, mod (b^x) m)| x <- [1..m-1]] of
        [] -> Nothing
        (x, _) : [] -> Just x


-- Probability of NO collision among t people
birthdayParadoxNoCollision t =
    foldl (\r i -> r * (1 - i/365)) 1 [1..t-1]


-- Number of input needed for a hash function collision with give probability
-- returns x where x is the exponent in 2^x
-- lambda: probability
-- n: output bit length
hashCollision lam n =
    (n+1)/2 + (logBase 2 $ sqrt $ log $ 1/(1-lam))

