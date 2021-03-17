#!/home/steed/Desktop/session_work/case_generate/introClass_cases/venv/bin/python3
import csv
from xml.dom.minidom import parse
from syllables.getcovrate import *

"""
    @定义全局变量
    @此文件是用于获取被测试文件的覆盖率信息矩阵的
"""

case_path = f"/home/steed/Desktop/session_work/case_generate/introClass_cases/{module_name}/tmpFile/cases"
matrix = []


def format_matrix(testpass, test_f_path):
    rootNode = parse(os.path.join(test_f_path, "test_cov.xml")).documentElement
    line_rate = rootNode.getAttribute("line-rate")  # 行覆盖率
    branch_rate = rootNode.getAttribute("branch-rate")  # 分支覆盖率
    lines_covered = rootNode.getAttribute("lines-covered")  # 覆盖行数
    branches_covered = rootNode.getAttribute("branches-covered")  # 覆盖分支数

    lineElements = rootNode.getElementsByTagName("line")
    hits_list = ["*"] * int(lineElements[-1].getAttribute("number"))
    hits_list.append("1" if testpass else "0")
    for line in lineElements:
        # print(line.getAttribute("number"), line.getAttribute("hits"))
        hits_list[int(line.getAttribute("number")) - 1] = line.getAttribute("hits")
    matrix.append(hits_list)


def add_title():
    title_list = [str(i) for i in range(1, len(matrix[0]))]
    title_list.append("pass?")  # used to check this case pass or not?
    matrix.insert(0, title_list)


def get_reporter():
    report_file = open("../tmpFile/report.csv", "w+", newline='', encoding="utf-8")
    writer = csv.writer(report_file)
    for row in matrix:
        writer.writerow(row)
    report_file.close()


def get_matrix():
    test_file_path = "/home/steed/Desktop/session_work/IntroClass/syllables/0cea42f9680f35f5a84c724c396d4d588b65c303453f9585562f2e2af8db74f5096a83a70b17c5126538222b111a0795a34e9fb6db95d62d771d01592abe3ff6/000"
    case_file = open(case_path, "r")
    rewrite_makefile(test_file_path)
    make_clean_test_file(test_file_path)
    for line in case_file.readlines():
        make_test_file(test_file_path)
        smokefile = open("./smokeFile.txt", "w+")
        smokefile.write(line)
        smokefile.seek(0)  # 这真的是让人咂舌
        color.warning("a case is:{}".format(line))
        test_pass = run_genprog_tests(
            test_exec_file_path=os.path.join(test_file_path, module_name),
            inp_data=smokefile,
            inp_data_str=line

        )

        smokefile.close()
        gcovr_save_test_xml(test_file_path)
        format_matrix(testpass=test_pass, test_f_path=test_file_path)
        make_clean_test_file(test_file_path)
    add_title()
    get_reporter()
    case_file.close()
    print("matrix:\n", matrix)


if __name__ == '__main__':
    get_matrix()