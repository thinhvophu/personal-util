import sys


def convert_format(file_path):
    with open(file_path, 'r') as file:
        input_text = file.read()
    output_text = []
    lines = input_text.split('\n')
    for line in lines:
        components = line.split()
        date = components[0]
        desc = ' '.join(components[1:-1])
        amount = components[-1]
        if amount.startswith('(') and amount.endswith(')'):
            amount = '-' + amount[1:-1]
        output_text.append(f"{date};{desc};{amount}")
    return '\n'.join(output_text)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python convert_format.py <file_path>")
    else:
        file_path = sys.argv[1]
        print(convert_format(file_path))
