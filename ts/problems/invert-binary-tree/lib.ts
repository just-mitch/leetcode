class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

export function invertTree(root: TreeNode | null): TreeNode | null {
  const queue: (TreeNode | null)[] = [];
  queue.push(root);
  while (queue.length > 0) {
    const node = queue.shift();
    if (node == null) {
      continue;
    }

    const temp = node.left;
    node.left = node.right;
    node.right = temp;

    queue.push(node.left);
    queue.push(node.right);
  }

  return root;
}
