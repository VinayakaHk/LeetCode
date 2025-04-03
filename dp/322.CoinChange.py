import math
def coinChange(coins,amount):
    print('coins ', coins)
    print('amount ', amount)
    dp = [math.inf] * (amount + 1)
    dp[0] = 0

    for a in range(1,amount + 1):
        print('a',a)
        for c in coins:
            print('c',c)
            if a - c >=0:
                print('a (',a,') - c (',c,') = ', a - c)
                tempamt = dp[a -c]
                print('dp[a - c]', tempamt)
                dp[a] = min(dp[a] , 1+tempamt)
                print('dp[a]', dp[a])
                print('min(dp[a], 1 + tempamt)', min(dp[a], 1 + tempamt))

            print('dp inner loop', dp)
            print('------------')
        print('dp outer loop', dp)
        print('---------------------------------------------------------------------------------------------------------------')
    return dp[amount] if dp[amount] != amount +1 else -1 


print('result : ')
print(coinChange([1,3,4,5] , 7))
