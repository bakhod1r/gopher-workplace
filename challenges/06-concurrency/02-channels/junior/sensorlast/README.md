# Last Reading

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

An IoT gateway streams readings from a device until the device disconnects.
The fleet dashboard shows the last reading received, or a no-data badge when
the device never reported.

## Task

Implement `LastReading` in [sensorlast.go](sensorlast.go) so that:

1. It drains `readings` until the gateway closes the stream.
2. It returns the most recently received reading and `true`.
3. It returns `0, false` when the device sent nothing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  LastReading(1, 2, 3)
Output: 3, true
```

**Example 2:**

```
Input:  LastReading() // closed, empty
Output: 0, false
```

**Example 3:**

```
Input:  LastReading(7)
Output: 7, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Drain and keep** | Overwrite a variable each iteration to end with the last value. |
| 2 | **"Seen" flag** | Distinguishes a real `0` reading from no data at all. |
| 3 | **`range` over a channel** | Ends at close, so "last" is well defined. |

## Hint

Keep two variables: the reading, overwritten each loop, and a bool set to
`true` the first time anything arrives.

## Validate

```bash
make verify
```
