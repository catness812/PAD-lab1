from multiprocessing import Process

def check_timeout(func):
    process = Process(target=func)
    process.start()
    process.join(timeout=10)

    if process.is_alive():
        process.terminate()
        return "Request timed out", 408
    else:
        return func()