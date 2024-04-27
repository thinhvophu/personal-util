import datetime


def convert_date(date_str):
    date_obj = datetime.datetime.strptime(date_str, '%d-%m-%Y')
    return date_obj.strftime('%d/%m')


def convert_format(input_file_path):
    with open(input_file_path, 'r') as input_file:
        for line in input_file:

            date = convert_date(line[:10])
            line = line[22:].replace('\n', '')
            descStart = 0
            if line[1] == '-':
                line = line[2:]
                descStart += 2
            components = line.split(' - ')
            if ' ' not in components[0]:
                amount = components[0]
            else:
                amount = components[0].split(' ')[-1]
            desc = line[descStart + len(components[0]):].replace(' - ', '')
            print(f"{date};{desc};{amount}")


if __name__ == "__main__":
    import sys

    if len(sys.argv) != 2:
        print("Usage: python convert_format.py <input_file_path>")
    else:
        input_file_path = sys.argv[1]
        convert_format(input_file_path)
